defmodule TopSecret do
  def to_ast(string) do
    {:ok, ast} = Code.string_to_quoted(string)
    ast
  end

  def decode_secret_message_part(ast, acc) do
    case ast do
      {f, _, [{:when, _, [{name, _, argument_list}, _]}, _]} when f in [:def, :defp] ->
        do_decode_secret_message_part(name, ast, argument_list, acc)

      {f, _, [{name, _, argument_list}, _]} when f in [:def, :defp] ->
        do_decode_secret_message_part(name, ast, argument_list, acc)

      _ ->
        {ast, acc}
    end
  end

  defp do_decode_secret_message_part(name, ast, argument_list, acc) do
    str = to_string(name)
    {ast, [cut_func_name(str, argument_list) | acc]}
  end

  defp cut_func_name(_str, list) when not is_list(list) or list == [], do: ""
  defp cut_func_name(str, list), do: String.slice(str, 0..(length(list) - 1))

  def decode_secret_message(string) do
    string
    |> to_ast()
    |> Macro.prewalk([], &decode_secret_message_part/2)
    |> then(fn {_, acc} -> acc end)
    |> Enum.reverse()
    |> Enum.join()
  end
end
