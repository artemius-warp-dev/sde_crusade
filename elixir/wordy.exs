defmodule Wordy do
  @operation_map %{
    "plus" => "+",
    "minus" => "-",
    "multiplied by" => "*",
    "divided by" => "/"
  }

  @doc """
  Calculate the math problem in the sentence.
  """

  @spec answer(String.t()) :: integer

  def answer("What is " <> question) do
    question
    |> de_wordify()
    |> String.split(" ", trim: true)
    |> Enum.map(&maybe_to_integer/1)
    |> calculate()
  end

  def answer(_), do: raise(ArgumentError)

  defp de_wordify(string) do
    Enum.reduce(@operation_map, string, fn {word, op}, s ->
      String.replace(s, word, op)
    end)
    |> String.trim_trailing("?")
  end

  defp calculate([number]), do: number

  defp calculate([x, "+", y | rest]) when is_integer(x) and is_integer(y),
    do: calculate([x + y | rest])

  defp calculate([x, "-", y | rest]) when is_integer(x) and is_integer(y),
    do: calculate([x - y | rest])

  defp calculate([x, "*", y | rest]) when is_integer(x) and is_integer(y),
    do: calculate([x * y | rest])

  defp calculate([x, "/", y | rest]) when is_integer(x) and is_integer(y),
    do: calculate([div(x, y) | rest])

  defp calculate(_), do: raise(ArgumentError)

  defp maybe_to_integer(s) do
    case Regex.match?(~r/^-?\d+$/, s) do
      true -> String.to_integer(s)
      false -> s
    end
  end
end

# defmodule Wordy do
#   @doc """
#   Calculate the math problem in the sentence.
#   """

#   @operations ["plus", "minus", "multiplied by", "divided by"]

#   @spec answer(String.t()) :: integer
#   def answer(question) do
#     question
#     |> parse_operations()
#     |> compute(0)
#   end

#   defp compute([], acc), do: acc

#   defp compute([operation | rest], acc) when operation in @operations do
#     [elem | tail] = rest
#     compute(tail, do_operation(operation, acc, elem))
#   end

#   defp compute([elem | rest], acc) do
#     compute(rest, elem)
#   end

#   defp do_operation(operation, left, right) do
#     case operation do
#       "plus" -> left + right
#       "minus" -> left - right
#       "multiplied by" -> left * right
#       "divided by" -> left / right
#     end
#   end

#   # defp validate_operations(list) do
#   #   cond do

#   #     list == [] ->
#   #       raise ArgumentError, "argument error"
#   #     not Enum.any?(list, fn elem -> elem in @operations end) ->
#   #       raise ArgumentError, "argument error"

#   #     true ->
#   #       list
#   #   end
#   # end

#   defp parse_operations(question) do
#     pattern =
#       ~r/(\d+)(?:\s+(plus|minus|multiplied by|divided by)\s+(\d+))?(?:\s+(plus|minus|multiplied by|divided by)\s+(\d+))*\??\z/

#     Regex.run(pattern, question, capture: :all_but_first)
#     |> IO.inspect()
#     |> then(fn
#       nil -> raise ArgumentError, "argument error"
#       res -> res
#     end)
#     |> List.flatten()
#     |> IO.inspect()
#     |> Enum.reduce({[], []}, fn
#       "-", {numbers, acc} ->
#         {["-" | numbers], acc}

#       operation, {numbers, acc} when operation in @operations ->
#         number = Enum.reverse(numbers) |> Enum.join()
#         # IO.inspect({[number, operation], acc})
#         {[], acc ++ [number, operation]}

#       str_number, {numbers, acc} ->
#         {[str_number | numbers], acc}
#     end)
#     # |> IO.inspect()
#     |> then(fn {str_digits, acc} ->
#       number = Enum.reverse(str_digits) |> Enum.join()
#       acc ++ [number]
#     end)
#     |> Enum.map(&parse_numbers/1)
#     |> IO.inspect()

#     # |> Enum.reduce({[], []}, fn
#     #   number, {numbers, acc} when is_integer(number) ->
#     #     {[number | numbers], acc}

#     #   operation, {numbers, acc} when is_binary(operation) ->
#     #     number = Enum.reverse(numbers) |> Integer.undigits()
#     #     {[], [number, operation] ++ acc}
#     # end)
#     # |> then(fn {numbers, acc} -> acc ++ numbers end)
#     # |> IO.inspect()
#   end

#   defp parse_numbers(str) do
#     case Integer.parse(str) do
#       {number, ""} -> number
#       _ -> str
#     end
#   end
# end

# "What is 52?"
# "What is 1 plus 5 minus -2?"
# "What is 3 plus 2 multiplied by 3?"
"What is 52 cubed?"
|> Wordy.answer()
|> IO.inspect()
