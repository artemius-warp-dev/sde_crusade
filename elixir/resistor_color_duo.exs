defmodule ResistorColorDuo do
  @doc """
  Calculate a resistance value from two colors
  """
  @spec value(colors :: [atom]) :: integer
  def value(colors) do
    colors
    |> Enum.take(2)
    |> Enum.map(&match_color/1)
    |> Enum.join("")
    |> IO.inspect()
    |> String.to_integer()
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

# defmodule ResistorColorDuo do
#   @doc """
#   Calculate a resistance value from two colors
#   """
#   @resistance_chart %{
#     :black => 0,
#     :brown => 1,
#     :red => 2,
#     :orange => 3,
#     :yellow => 4,
#     :green => 5,
#     :blue => 6,
#     :violet => 7,
#     :grey => 8,
#     :white => 9
#   }

#   @spec value(colors :: [atom]) :: integer
#   def value(colors) do
#     colors
#     |> Enum.map(&(@resistance_chart[&1]))
#     |> Enum.take(2)
#     |> Integer.undigits
#  end
# end
