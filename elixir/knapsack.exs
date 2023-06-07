defmodule Knapsack do
  @doc """
  Return the maximum value that a knapsack can carry.
  """
  @spec maximum_value(items :: [%{value: integer, weight: integer}], maximum_weight :: integer) ::
          integer
  def maximum_value(items, maximum_weight) do
    number_of_items = length(items)
    dp = get_cache(number_of_items, maximum_weight)
    res = maximum_value(items, maximum_weight, dp)
    res[{number_of_items, maximum_weight}]
  end

  defp maximum_value(items, maximum_weight, dp) do
    Enum.with_index(items, 1)
    |> Enum.reduce(dp, fn item, dp ->
      fill_cache(item, 1, maximum_weight, items, dp)
    end)
  end

  defp fill_cache({item, i}, weight, maximum_weight, items, dp) when weight <= maximum_weight do
    if pack?({item, i}, weight) do
      new_value =
        Enum.max([
          dp[{i - 1, weight}],
          dp[{i - 1, weight - item.weight}] + item.value
        ])

      fill_cache(
        {item, i},
        weight + 1,
        maximum_weight,
        items,
        Map.put(dp, {i, weight}, new_value)
      )
    else
      fill_cache(
        {item, i},
        weight + 1,
        maximum_weight,
        items,
        Map.put(dp, {i, weight}, dp[{i - 1, weight}])
      )
    end
  end

  defp fill_cache(_, _, _, _, dp), do: dp

  defp pack?({item, _i}, capacity) do
    item.weight <= capacity
  end

  defp get_cache(items_amount, maximum_weight) do
    for w <- 0..maximum_weight,
        i <- 0..items_amount,
        into: %{},
        do: {{i, w}, 0}
  end
end

# def maximum_value(items, maximum_weight) do
#   do_max_val(items, maximum_weight, 0)
# end

# defp do_max_val([%{value: value, weight: weight} | rest], capacity, acc) do
#   if weight > capacity do
#     acc
#   else
#     with_this = do_max_val(rest, capacity - weight, acc + value)

#     without_this = do_max_val(rest, capacity, acc)

#     max(with_this, without_this)
#   end
# end

# defp do_max_val([], _, acc), do: acc

# defp do_max_val(_, 0, acc), do: acc

items = [
  %{value: 350, weight: 25},
  %{value: 400, weight: 35},
  %{value: 450, weight: 45},
  %{value: 20, weight: 5},
  %{value: 70, weight: 25},
  %{value: 8, weight: 3},
  %{value: 5, weight: 2},
  %{value: 5, weight: 2}
]

maximum_weight = 104
# assert Knapsack.maximum_value(items, maximum_weight) == 900

# items = [
#   %{value: 5, weight: 2},
#   %{value: 5, weight: 2},
#   %{value: 5, weight: 2},
#   %{value: 5, weight: 2},
#   %{value: 21, weight: 10}
# ]

# maximum_weight = 10
# assert Knapsack.maximum_value(items, maximum_weight) == 21

# items = [
#   %{value: 20, weight: 2},
#   %{value: 20, weight: 2},
#   %{value: 20, weight: 2},
#   %{value: 20, weight: 2},
#   %{value: 50, weight: 10}
# ]

# maximum_weight = 10

# assert Knapsack.maximum_value(items, maximum_weight) == 80

Knapsack.maximum_value(items, maximum_weight) |> IO.inspect()
