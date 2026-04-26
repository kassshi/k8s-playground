import { ConnectError } from "@connectrpc/connect";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { CreateTodo } from "../api/todos";

type CreateTodoFormValues = {
  title: string;
  decription: string;
};

export function CreateTodoModal() {
  const [isOpen, setIsOpen] = useState(false);
  const { register, handleSubmit } = useForm<CreateTodoFormValues>();
  const onSubmit = async (data: CreateTodoFormValues) => {
    try {
      await CreateTodo(data.title, data.decription);
      setIsOpen(false);
    } catch (error) {
      if (error instanceof ConnectError) {
        console.log("Create todo failed:", error.message);
      }
    }
  };
  return (
    <>
      <button onClick={() => setIsOpen(true)}>Create Todo</button>
      {isOpen && (
        <div className="modal">
          <div className="modal-content">
            <h2>Create Todo</h2>
            <form onSubmit={handleSubmit(onSubmit)}>
              <input
                type="text"
                placeholder="Title"
                {...register("title")}
                required
              />
              <textarea
                placeholder="Description"
                {...register("decription")}
                required
              />
              <button type="submit">Create</button>
            </form>
          </div>
        </div>
      )}
    </>
  );
}
