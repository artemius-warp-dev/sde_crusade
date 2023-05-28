defmodule Change do
  @doc """
    Determine the least number of coins to be given to the user such
    that the sum of the coins' value would equal the correct amount of change.
    It returns {:error, "cannot change"} if it is not possible to compute the
    right amount of coins. Otherwise returns the tuple {:ok, list_of_coins}

    ## Examples

      iex> Change.generate([5, 10, 15], 3)
      {:error, "cannot change"}

      iex> Change.generate([1, 5, 10], 18)
      {:ok, [1, 1, 1, 5, 10]}

  """

  @spec generate(list, integer) :: {:ok, list} | {:error, String.t()}
  def generate(coins, target) do
    dp = for x <- 0..target, into: %{}, do: {x, []}

    1..target
    |> Enum.reduce(dp, &changes_for(&1, &2, coins))
    |> then(fn acc ->
      case acc[target] do
        [] -> {:error, "cannot change"}
        change -> {:ok, change}
      end
    end)
  end

  defp changes_for(target, acc, coins) do
    change = change(coins, target, acc)
    Map.put(acc, target, change)
  end

  defp change(coins, target, acc) do
    coins
    |> Enum.filter(fn x -> x <= target end)
    |> Enum.map(fn x ->
      if acc[target - x] != [] or target - x == 0 do
        [x | acc[target - x]] |> List.flatten()
      else
        []
      end
    end)
    |> Enum.filter(fn x ->
      length(x) > 0
    end)
    |> Enum.min_by(&length/1, fn -> [] end)
  end
end

Change.generate([1, 5, 10], 18) |> IO.inspect(label: "Result", charlists: :as_list)
