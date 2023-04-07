defmodule NucleotideCount do
  @nucleotides [?A, ?C, ?G, ?T]

  @doc """
  Counts individual nucleotides in a DNA strand.

  ## Examples

  iex> NucleotideCount.count('AATAA', ?A)
  4

  iex> NucleotideCount.count('AATAA', ?T)
  1
  """
  @spec count(charlist(), char()) :: non_neg_integer()
  def count(strand, nucleotide) when nucleotide in @nucleotides do
    strand
    |> Enum.filter(&(&1 == nucleotide))
    |> Enum.count()
  end

  @doc """
  Returns a summary of counts by nucleotide.

  ## Examples

  iex> NucleotideCount.histogram('AATAA')
  %{?A => 4, ?T => 1, ?C => 0, ?G => 0}
  """
  @spec histogram(charlist()) :: map()
  def histogram(strand) do
    map = @nucleotides |> Map.new(fn x -> {"?#{[x]}", 0} end)

    strand
    |> Enum.reduce(map, fn x, acc ->
      Map.put(acc, "?#{[x]}", count(strand, x))
    end)
  end
end

# NucleotideCount.count('AATAA', ?C) |> IO.inspect()
NucleotideCount.histogram('AATAA') |> IO.inspect()
