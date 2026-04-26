import { createClient } from "@connectrpc/connect";
import { Status, TodoService } from "../../../gen/todo/v1/todo_pb";
import { transport } from "../../../shared/lib/connect";
const client = createClient(TodoService, transport);

export async function ListTodos() {
  return await client.listTodos({});
}

export async function CreateTodo(title: string, description: string) {
  return await client.createTodo({
    todo: {
      title: title,
      description: description,
      status: Status.TODO,
    },
  });
}
