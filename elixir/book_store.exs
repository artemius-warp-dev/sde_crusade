defmodule BookStore do
  @typedoc "A book is represented by its number in the 5-book series"
  @type book :: 1 | 2 | 3 | 4 | 5

  @doc """
  Calculate lowest price (in cents) for a shopping basket containing books.
  """
  @spec total(basket :: [book]) :: integer
  def total(basket) do
    groups = Enum.frequencies(basket)
    sorted = Enum.sort_by(basket, &Map.get(groups, &1), :desc)

    2..5
    |> Enum.map(&group(sorted, &1))
    |> Enum.map(&calculate/1)
    |> Enum.min()
  end

  defp group(list, size, group \\ [], acc \\ [])
  defp group([], _size, group, acc), do: [group | acc]

  defp group(list, size, group, acc) when length(group) < size do
    case Enum.find(list, fn e -> e not in group end) do
      nil ->
        group(list, size, [], [group | acc])

      elem ->
        updated_list = List.delete(list, elem)

        group(
          updated_list,
          size,
          [elem | group],
          acc
        )
    end
  end

  defp group(list, size, group, acc), do: group(list, size, [], [group | acc])

  defp calculate(books) when is_integer(books) do
    books * 800
  end

  defp calculate(list_of_combinations) do
    list_of_combinations
    |> Enum.reduce(0, fn combination, acc ->
      books = length(combination)
      acc + apply_discount(calculate(books), calculate_discount(books))
    end)
  end

  defp calculate_discount(num_unique_books) do
    case num_unique_books do
      2 -> 0.05
      3 -> 0.10
      4 -> 0.20
      5 -> 0.25
      _ -> 0
    end
  end

  defp apply_discount(cost, discount) do
    (cost * (1 - discount)) |> trunc()
  end
end

# Example usage
# basket = [1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5]

# 1, 1, 2, 2, 3, 3, 4, 5
basket = [1, 1, 2, 3, 4]
# [1] [1 2 3 4]
BookStore.total(basket) |> IO.inspect()
