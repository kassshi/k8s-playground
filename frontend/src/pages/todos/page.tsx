import { useQuery } from "@tanstack/react-query";
import { ListTodos } from "../../features/todo/api/todos";
import { CreateTodoModal } from "../../features/todo/components/createTodoModal";
import type { Todo } from "../../gen/todo/v1/todo_pb";
export function TodosPage() {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ["todos"],
    queryFn: ListTodos,
  });

  if (isPending) {
    return <span>Loading...</span>;
  }

  if (isError) {
    return <span>Error: {error.message}</span>;
  }

  return (
    <div>
      <CreateTodoModal />
      <ul>
        {data.todos.map((todo: Todo) => (
          <li key={todo.name}>{todo.title}</li>
        ))}
      </ul>
    </div>
  );
}
