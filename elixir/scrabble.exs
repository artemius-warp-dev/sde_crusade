defmodule Scrabble do
  @doc """
  Calculate the scrabble score for the word.
  """
  @spec score(String.t()) :: non_neg_integer
  def score(word) do
    word
    |> String.upcase()
    |> String.codepoints()
    |> Enum.reduce(0, &calculate/2)
  end

  defp1 calculate(ch, acc) do
    cond do
      ch in ["A", "E", "I", "O", "U", "L", "N", "R", "S", "T"] -> acc + 1
      ch in ["D", "G"] -> acc + 2
      ch in ["B", "C", "M", "P"] -> acc + 3
      ch in ["F", "H", "V", "W", "Y"] -> acc + 4
      ch in ["K"] -> acc + 5
      ch in ["J", "X"] -> acc + 8
      ch in ["Q", "Z"] -> acc + 10
      true -> acc
    end
  end
end
