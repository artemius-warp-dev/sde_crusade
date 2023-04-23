defmodule ResistorColorTrio do
  @doc """
  Calculate the resistance value in ohms from resistor colors
  """
  @spec label(colors :: [atom]) :: {number, :ohms | :kiloohms | :megaohms | :gigaohms}
  def label(colors) do
    colors
    |> Enum.take(3)
    |> value()
    |> then(fn {v, zeros} -> 
      {rem, prefix} = get_metric_prefix(zeros)
      {"#{v}#{String.duplicate("0", rem)}" |> String.to_integer(), prefix}
    end)
    
  end

  defp get_metric_prefix(value) do
    cond do
      value == 0 -> {0, :ohms}
      value < 3 -> {value, :ohms}
      value == 3 -> {0, :kiloohms}
      value < 6 -> {value - 3, :kiloohms}
      value == 6 -> {0, :megaohms}
      value < 9 ->  {value - 6, :megaohms}
      value == 9 -> {0, :gigaohms}
      value > 9 -> {value - 9, :gigaohms}
    end
  end


  defp value(colors) do
    colors
    |> Enum.map(&match_color/1)
    |> Enum.join("")
    |> String.split_at(2)
    |> then(fn {value, zeros} -> "#{value}#{String.duplicate("0", String.to_integer(zeros))}" end)
    |> String.graphemes()
    |> Enum.reduce({"", 0}, fn
        "0", {"", zeros} -> {"0", zeros}
        "0", {x, zeros} -> {x, zeros + 1}
        v, {x, zeros} -> {"#{x}#{v}", zeros}
        end)
  end

  defp match_color(color) do
    case color do
      :black -> 0
      :brown -> 1
      :red -> 2
      :orange -> 3
      :yellow -> 4
      :green -> 5
      :blue -> 6
      :violet -> 7
      :grey -> 8
      :white -> 9
    end
  end

end
