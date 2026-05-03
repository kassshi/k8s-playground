import { Field, FieldGroup } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { ConnectError } from "@connectrpc/connect";
import { useState } from "react";

import { useQueryClient } from "@tanstack/react-query";
import { PlusIcon } from "lucide-react";
import { useForm } from "react-hook-form";
import { Button } from "../../../components/ui/button";
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "../../../components/ui/dialog";
import { Label } from "../../../components/ui/label";
import { CreateTodo } from "../api/todos";
type CreateTodoFormValues = {
  title: string;
  decription: string;
};
export function CreateTodoModal() {
  const { register, handleSubmit, reset } = useForm<CreateTodoFormValues>();
  const queryClient = useQueryClient();
  const [open, setOpen] = useState(false);
  const onSubmit = async (data: CreateTodoFormValues) => {
    try {
      await CreateTodo(data.title, data.decription);
      await queryClient.invalidateQueries({ queryKey: ["todos"] });
      reset();
      setOpen(false);
    } catch (error) {
      if (error instanceof ConnectError) {
        console.log("Create todo failed:", error.message);
      }
    }
  };
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button className="mt-4">
          <PlusIcon />
          Create Todo
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <form onSubmit={handleSubmit(onSubmit)}>
          <DialogHeader>
            <DialogTitle>Create Todo</DialogTitle>
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
          </FieldGroup>
          <DialogFooter className="mt-4">
            <DialogClose asChild>
              <Button variant="outline">Cancel</Button>
            </DialogClose>
            <Button type="submit">Create</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
