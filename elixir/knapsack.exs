defmodule Knapsack do
  @doc """
  Return the maximum value that a knapsack can carry.
  """
  @spec maximum_value(items :: [%{value: integer, weight: integer}], maximum_weight :: integer) ::
          integer

  def maximum_value(items, maximum_weight) do
    items
    |> Enum.sort_by(&({&1.weight, &1.value}), fn {w1,v1},{w2,v2} ->
      IO.inspect({{w1,v1},{w2, v2}})
       w1 <= w2
    end)
    |> IO.inspect()
    |> do_greedy_scan(maximum_weight)
  end


  defp do_greedy_scan(items, maximum_weight) do
    items
    |> Enum.reduce({0, maximum_weight}, fn
      x, {profit, m_weight} when x.weight <= m_weight ->
        {x.value+profit, m_weight - x.weight}
      _, acc -> acc
    end)
  end
  
  # defp scan([], _), do: []
  # defp scan(items, maximum_weight) do
  #   Enum.map(items, fn item -> sum(item.weight, items -- [item], maximum_weight, item.value) end)
  # end

  # defp sum(_, [], _, acc), do: acc

  # defp sum(weight, [h | tail], maximum_weight, acc) do
  #   if weight + h.weight <= maximum_weight do
  #     sum(weight + h.weight, tail, maximum_weight, acc + h.value)
  #   else
  #     sum(weight, tail, maximum_weight, acc)
  #   end
  # end
end

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
#assert Knapsack.maximum_value(items, maximum_weight) == 900

Knapsack.maximum_value(items, maximum_weight) |> IO.inspect()
