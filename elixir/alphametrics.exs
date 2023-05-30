defmodule Alphametics do
  @type puzzle :: binary
  @type solution :: %{required(?A..?Z) => 0..9}

  @doc """
  Takes an alphametics puzzle and returns a solution where every letter
  replaced by its number will make a valid equation. Returns `nil` when
  there is no valid solution to the given puzzle.

  ## Examples

    iex> Alphametics.solve("I + BB == ILL")
    %{?I => 1, ?B => 9, ?L => 0}

    iex> Alphametics.solve("A == B")
    nil
  """
  @spec solve(puzzle) :: solution | nil
  def solve(puzzle) do
    {struct, letters} = parse(puzzle)

    find_solution(
      struct,
      letters,
      Enum.to_list(0..9),
      %{}
    )
  end

  defp find_solution(struct, [], _, acc), do: check_solution(struct, acc)

  defp find_solution(struct, letters, digits, acc) do
    Enum.find_value(
      digits,
      &do_solve(struct, letters, &1, digits -- [&1], acc)
    )
  end

  defp do_solve(struct, [ch | rest], digit, digits, acc) do
    # cond do
    #  ch in struct.nonzeros ->
    #    false

    #  digit > 0 or ch not in struct.nonzeros ->
    find_solution(struct, rest, digits, Map.put(acc, ch, digit))
    # end
  end

  defp do_solve(struct, [], _, _, acc), do: check_solution(struct, acc)

  defp check_solution(struct, acc) do
    sum =
      struct.addends
      |> Enum.map(&to_number(&1, acc))
      |> Enum.reduce(&Kernel.+/2)
      |> Kernel.==(to_number(struct.sum, acc))
      |> then(fn
        true -> acc
        _ -> false
      end)
  end

  defp to_number(chars, acc) do
    chars
    |> Enum.map(&acc[&1])
    |> Integer.undigits()
  end

  defp parse(puzzle) do
    words =
      puzzle
      |> String.split(" ")
      |> Enum.filter(fn str -> String.match?(str, ~r/^[A-Z]+$/) end)
      |> Enum.map(&String.to_charlist/1)
      |> Enum.reverse()

    [sum | addends] = words

    nonzeros =
      words
      |> Enum.map(&hd/1)
      |> Enum.uniq()

    letters =
      words
      |> List.flatten()
      |> Enum.uniq()

    {%{addends: addends, sum: sum, nonzeros: nonzeros}, letters}
  end

  defp solve(left, right, result, max_sum, acc) do
  end
end

Alphametics.solve("SEND + MORE == MONEY") |> IO.inspect()
# 21 + 81 == 102
# 21 + 91 == 112
