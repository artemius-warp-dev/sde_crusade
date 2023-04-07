defmodule Pangram do
  @doc """
  Determines if a word or sentence is a pangram.
  A pangram is a sentence using every letter of the alphabet at least once.

  Returns a boolean.

    ## Examples

      iex> Pangram.pangram?("the quick brown fox jumps over the lazy dog")
      true

  """

  @cache Map.new(?a..?z, &{"#{[&1]}", 0})

  @spec pangram?(String.t()) :: boolean
  def pangram?(sentence) do
    sentence
    |> String.downcase()
    |> String.graphemes()
    |> Enum.reduce(@cache, &calculate(&1, &2, Map.has_key?(&2, &1)))
    |> Enum.all?(fn {_, x} -> x > 0 end)
  end

 

  defp calculate(x, state, exist?) when exist?, do: Map.update!(state, x, &(&1 + 1))

  defp calculate(_, state, _), do: state
end

Pangram.pangram?("the quick brown fox jumps over the lazy dog") |> IO.inspect()
