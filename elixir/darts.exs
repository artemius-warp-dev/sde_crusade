defmodule Darts do
  @type position :: {number, number}

  @doc """
  Calculate the score of a single dart hitting a target
  """
  @spec score(position) :: integer
  def score({x, y}) do
    [{x, y}, {0, 0}]
    |> get_distance()
    |> do_score()
  end

  defp do_score(distance) when distance > 10, do: 0
  defp do_score(distance) when distance <= 10 and distance > 5, do: 1
  defp do_score(distance) when distance <= 5 and distance > 1, do: 5
  defp do_score(distance) when distance <= 1 and distance >= 0, do: 10

  defp get_distance([{x1, y1}, {x2, y2}]) do
    :math.sqrt(abs(x1 - x2) ** 2 + abs(y1 - y2) ** 2)
  end
end

Darts.score({-9, 9}) |> IO.inspect()
