import { Badge } from "@/components/ui/badge";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { useState } from "react";
import { Status, type Todo } from "../../../gen/todo/v1/todo_pb";
import { EditTodoModal } from "./editTodoModal";

export function TodoItem({ todo }: { todo: Todo }) {
  const [openEditModal, setOpenEditModal] = useState(false);
  return (
    <div>
      <EditTodoModal
        todo={todo}
        open={openEditModal}
        onOpenChange={setOpenEditModal}
      />
      <Card
        className="mb-4 cursor-pointer"
        onClick={() => setOpenEditModal(true)}
      >
        <CardHeader>
          {" "}
          <CardTitle>{todo.title}</CardTitle>
        </CardHeader>
        <CardContent>
          <p
            className="text-sm text-muted-foreground
  mb-3"
          >
            {todo.description}
          </p>
          {todo.status == Status.TODO && (
            <Badge className="bg-sky-50 text-sky-700 dark:bg-sky-950 dark:text-sky-300">
              TODO
            </Badge>
          )}
          {todo.status == Status.IN_PROGRESS && (
            <Badge className="bg-yellow-50 text-yellow-700 dark:bg-yellow-950 dark:text-yellow-300">
              IN_PROGRESS
            </Badge>
          )}
          {todo.status == Status.DONE && (
            <Badge className="bg-green-50 text-green-700 dark:bg-green-950 dark:text-green-300">
              DONE
            </Badge>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
