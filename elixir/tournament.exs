defmodule Tournament do
  @doc """
  Given `input` lines representing two teams and whether the first of them won,
  lost, or reached a draw, separated by semicolons, calculate the statistics
  for each team's number of games played, won, drawn, lost, and total points
  for the season, and return a nicely-formatted string table.

  A win earns a team 3 points, a draw earns 1 point, and a loss earns nothing.

  Order the outcome by most total points for the season, and settle ties by
  listing the teams in alphabetical order.
  """
  @spec tally(input :: list(String.t())) :: String.t()
  def tally(input) do
    input
    |> parse(%{})
    |> Map.to_list()
    |> Enum.sort(fn {name_a, {_, _, _, _, p_a}}, {name_b, {_, _, _, _, p_b}} ->
      cond do
        p_a == p_b -> name_a < name_b
        true -> p_a > p_b
      end
    end)
    |> Enum.map(&add_line/1)
    |> then(fn list -> [add_head() | list] end)
    |> Enum.join("\n")
  end

  defp parse([], acc), do: acc

  defp parse([h | t], acc) do
    parse(t, calculate(h, acc))
  end

  defp calculate(str, acc) do
    case String.split(str, ";") do
      [team_a, team_b, "win"] -> update_result(acc, team_a, team_b, :win)
      [team_a, team_b, "draw"] -> update_result(acc, team_a, team_b, :draw)
      [team_a, team_b, "loss"] -> update_result(acc, team_a, team_b, :loss)
      _ -> acc
    end
  end

  defp update_result(acc, team_a, team_b, :win) do
    acc
    |> Map.update(team_a, {1, 1, 0, 0, 3}, fn {mp, w, d, l, s} -> {mp + 1, w + 1, d, l, s + 3} end)
    |> Map.update(team_b, {1, 0, 0, 1, 0}, fn {mp, w, d, l, s} -> {mp + 1, w, d, l + 1, s} end)
  end

  defp update_result(acc, team_a, team_b, :draw) do
    acc
    |> Map.update(team_a, {1, 0, 1, 0, 1}, fn {mp, w, d, l, s} -> {mp + 1, w, d + 1, l, s + 1} end)
    |> Map.update(team_b, {1, 0, 1, 0, 1}, fn {mp, w, d, l, s} -> {mp + 1, w, d + 1, l, s + 1} end)
  end

  defp update_result(acc, team_a, team_b, :loss) do
    acc
    |> Map.update(team_a, {1, 0, 0, 1, 0}, fn {mp, w, d, l, s} -> {mp + 1, w, d, l + 1, s} end)
    |> Map.update(team_b, {1, 1, 0, 0, 3}, fn {mp, w, d, l, s} -> {mp + 1, w + 1, d, l, s + 3} end)
  end

  defp add_head() do
    line = "| MP |  W |  D |  L |  P"
    "#{String.pad_trailing("Team", 30)} #{line}"
  end

  defp add_line({team, {mp, w, d, l, s}}) do
    line =
      "| #{String.pad_leading("#{mp}", 2)} | #{String.pad_leading("#{w}", 2)} | #{String.pad_leading("#{d}", 2)} | #{String.pad_leading("#{l}", 2)} | #{String.pad_leading("#{s}", 2)}"

    "#{String.pad_trailing(team, 30)} #{line}"
  end
end

input = [
  "Allegoric Alaskans;Blithering Badgers;win"
]

Tournament.tally(input) |> IO.inspect()
