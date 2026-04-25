import { createClient } from "@connectrpc/connect";
import { TodoService } from "../../../gen/todo/v1/todo_pb";
import { transport } from "../../../shared/lib/connect";
const client = createClient(TodoService, transport);

export async function listTodos() {
  return await client.listTodos({});
}
