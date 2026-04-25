import { useQuery } from "@tanstack/react-query";
import { listTodos } from "../../features/todo/api/todos";
import type { Todo } from "../../gen/todo/v1/todo_pb";
export function TodosPage() {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ["todos"],
    queryFn: listTodos,
  });

  if (isPending) {
    return <span>Loading...</span>;
  }

  if (isError) {
    return <span>Error: {error.message}</span>;
  }

  return (
    <ul>
      {data.todos.map((todo: Todo) => (
        <li key={todo.name}>{todo.title}</li>
      ))}
    </ul>
  );
}
