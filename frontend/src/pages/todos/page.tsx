import { TodoItem } from "@/features/todo/components/todoItem";
import { useQuery } from "@tanstack/react-query";
import { ListTodos } from "../../features/todo/api/todos";
import { CreateTodoModal } from "../../features/todo/components/createTodoModal";
import { Status, type Todo } from "../../gen/todo/v1/todo_pb";
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

  const todoItems = data.todos.filter((todo) => todo.status == Status.TODO);

  const inProgessItems = data.todos.filter(
    (todo) => todo.status == Status.IN_PROGRESS,
  );

  const doneItems = data.todos.filter((todo) => todo.status == Status.DONE);
  return (
    <div>
      <CreateTodoModal />

      <div className="grid grid-cols-3 gap-4 mt-4">
        <div>
          <h2 className="text-lg font-bold mb-4">TODO</h2>
          {todoItems.map((todo: Todo) => (
            <TodoItem todo={todo} />
          ))}
        </div>

        <div>
          <h2 className="text-lg font-bold mb-4">IN PROGRESS</h2>
          {inProgessItems.map((todo: Todo) => (
            <TodoItem todo={todo} />
          ))}
        </div>

        <div>
          <h2 className="text-lg font-bold mb-4">DONE</h2>
          {doneItems.map((todo: Todo) => (
            <TodoItem todo={todo} />
          ))}
        </div>
      </div>
    </div>
  );
}
