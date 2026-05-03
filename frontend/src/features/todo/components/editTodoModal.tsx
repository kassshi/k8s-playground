import { Button } from "@/components/ui/button";

import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogFooter,
    DialogHeader,
    DialogTitle,
} from "@/components/ui/dialog";
import { Field, FieldGroup } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { Status, type Todo } from "@/gen/todo/v1/todo_pb";
import { ConnectError } from "@connectrpc/connect";
import { useQueryClient } from "@tanstack/react-query";
import { Controller, useForm } from "react-hook-form";
import { Label } from "../../../components/ui/label";
import { DeleteTodo, UpdateTodo } from "../api/todos";
type EditTodoFormValues = {
  name: string;
  title: string;
  decription: string;
  status: Status;
};
export function EditTodoModal({
  todo,
  open,
  onOpenChange,
}: {
  todo: Todo;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}) {
  const { register, handleSubmit, reset, control } =
    useForm<EditTodoFormValues>({
      defaultValues: {
        name: todo.name,
        title: todo.title,
        decription: todo.description,
        status: todo.status,
      },
    });

  const queryClient = useQueryClient();
  const onSubmit = async (data: EditTodoFormValues) => {
    try {
      await UpdateTodo(data.name, data.title, data.decription, data.status);
      await queryClient.invalidateQueries({ queryKey: ["todos"] });
      reset();
      onOpenChange(false);
    } catch (error) {
      if (error instanceof ConnectError) {
        console.log("Edit todo failed:", error.message);
      }
    }
  };

  const deleteTodo = async (name: string) => {
    try {
      await DeleteTodo(name);
      await queryClient.invalidateQueries({ queryKey: ["todos"] });
      reset();
      onOpenChange(false);
    } catch (error) {
      if (error instanceof ConnectError) {
        console.log("Edit todo failed:", error.message);
      }
    }
  };
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-sm">
        <form onSubmit={handleSubmit(onSubmit)}>
          <DialogHeader>
            <DialogTitle>Edit Todo</DialogTitle>
          </DialogHeader>
          <FieldGroup>
            <Field>
              <Label htmlFor="title">Title</Label>
              <Input
                type="text"
                id="title"
                placeholder="Title"
                {...register("title")}
                required
              />
            </Field>
            <Field>
              <Label htmlFor="description">Description</Label>
              <Textarea
                id="description"
                placeholder="Description"
                {...register("decription")}
                required
              />
            </Field>
            <Field>
              <Controller
                control={control}
                name="status"
                render={({ field }) => (
                  <Select
                    onValueChange={(value) => field.onChange(Number(value))}
                    defaultValue={String(field.value)}
                  >
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectItem value={String(Status.TODO)}>
                          TODO
                        </SelectItem>
                        <SelectItem value={String(Status.IN_PROGRESS)}>
                          IN PROGRESS
                        </SelectItem>
                        <SelectItem value={String(Status.DONE)}>
                          DONE
                        </SelectItem>
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                )}
              />
            </Field>
          </FieldGroup>
          <DialogFooter className="mt-4">
            <Button variant="destructive" onClick={() => deleteTodo(todo.name)}>
              {" "}
              Delete{" "}
            </Button>
            <DialogClose asChild>
              <Button variant="outline">Cancel</Button>
            </DialogClose>
            <Button type="submit">Edit</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
