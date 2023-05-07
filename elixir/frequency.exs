defmodule Frequency do
  @doc """
  Count letter frequency in parallel.

  Returns a map of characters to frequencies.

  The number of worker processes to use can be set with 'workers'.
  """
  @spec frequency([String.t()], pos_integer) :: map
  def frequency(texts, workers) do
    # 1 split to part equal to workers
    # 2 split partions to tasks
    # 3 calculate results
    texts
    |> Enum.chunk_every(workers)
    |> Enum.map(fn x -> Task.async(fn -> calculate(x) end) end)
    |> Task.await_many()
    |> Enum.reduce(%{}, fn x, acc -> merge(acc, x) end)
  end

  defp merge(map1, map2) do
    Map.merge(map1, map2, fn _k, v1, v2 -> v1 + v2 end) |> IO.inspect()
  end

  defp calculate(strings) do
    strings
    |> Enum.reduce(%{}, fn line, acc ->
      line
      |> String.replace(~r/[^[:alpha:]]/u, "")
      |> String.downcase()
      |> String.graphemes()
      |> Enum.reduce(%{}, fn ch, acc -> Map.update(acc, ch, 1, fn x -> x + 1 end) end)
      |> merge(acc)
    end)
  end
end
Frequency.calculate(["fDcfddd", "cCCCCeDDD"]) |> IO.inspect()
