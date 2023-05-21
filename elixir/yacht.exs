defmodule Yacht do
  @type category ::
          :ones
          | :twos
          | :threes
          | :fours
          | :fives
          | :sixes
          | :full_house
          | :four_of_a_kind
          | :little_straight
          | :big_straight
          | :choice
          | :yacht

  @doc """
  Calculate the score of 5 dice using the given category's scoring method.
  """
  @spec score(category :: category(), dice :: [integer]) :: integer
  def score(categor, dice)

  def score(:yacht, dices) do
    if Enum.all?(dices, fn x -> x == List.first(dices) end) do
      50
    else
      0
    end
  end

  def score(:ones, dices), do: Enum.count(dices, fn x -> x == 1 end) * 1
  def score(:twos, dices), do: Enum.count(dices, fn x -> x == 2 end) * 2
  def score(:threes, dices), do: Enum.count(dices, fn x -> x == 3 end) * 3
  def score(:fours, dices), do: Enum.count(dices, fn x -> x == 4 end) * 4
  def score(:fives, dices), do: Enum.count(dices, fn x -> x == 5 end) * 5
  def score(:sixes, dices), do: Enum.count(dices, fn x -> x == 6 end) * 6

  def score(:full_house, dices) do
    grouped =
      dices
      |> Enum.reduce(%{}, fn x, acc -> Map.update(acc, x, 1, fn v -> v + 1 end) end)
      |> Map.filter(fn {_k, v} -> v == 3 or v == 2 end)

    if Map.values(grouped) |> Enum.sort() == [2, 3] do
      Enum.reduce(grouped, 0, fn {k, v}, acc -> acc + v * k end)
    else
      0
    end
  end

  def score(:four_of_a_kind, dices) do
    grouped =
      dices
      |> Enum.reduce(%{}, fn x, acc ->
        Map.update(acc, x, 1, fn
          v when v < 4 -> v + 1
          v -> v
        end)
      end)
      |> Map.filter(fn {_k, v} -> v >= 4 end)

    if Map.values(grouped) == [4] do
      Enum.reduce(grouped, 0, fn {k, v}, acc -> acc + v * k end)
    else
      0
    end
  end

  def score(:little_straight, dices) do
    dices
    |> Enum.sort()
    |> then(fn
      [1, 2, 3, 4, 5] -> 30
      _ -> 0
    end)
  end

  def score(:big_straight, dices) do
    dices
    |> Enum.sort()
    |> then(fn
      [2, 3, 4, 5, 6] -> 30
      _ -> 0
    end)
  end

  def score(:choice, dices), do: Enum.sum(dices)
end
