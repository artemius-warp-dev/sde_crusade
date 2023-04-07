defmodule Hamming do
  @doc """
  Returns number of differences between two strands of DNA, known as the Hamming Distance.

  ## Examples

  iex> Hamming.hamming_distance('AAGTCATA', 'TAGCGATC')
  {:ok, 4}
  """
  @spec hamming_distance([char], [char]) :: {:ok, non_neg_integer} | {:error, String.t()}
  def hamming_distance(strand1, strand2)

  def hamming_distance(s1, s2) when length(s1) != length(s2),
    do: {:error, "strands must be of equal length"}

  def hamming_distance(strand1, strand2) do
    diff_count =
      0..length(strand1)
      |> Enum.filter(&(Enum.at(strand1, &1) != Enum.at(strand2, &1)))
      |> Enum.count()

    {:ok, diff_count}
  end
end

Hamming.hamming_distance('GGACGGATTCTG', 'AGGACGGATTCT') |> IO.inspect()
