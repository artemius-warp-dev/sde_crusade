defmodule Anagram do
  @doc """
  Returns all candidates that are anagrams of, but not equal to, 'base'.
  """
  @spec match(String.t(), [String.t()]) :: [String.t()]
  def match(base, candidates) do
    base_list = base |> String.graphemes()

    candidates
    |> Enum.map(&String.graphemes/1)
    |> Enum.filter(&(not case_match?(&1, base_list)))
    |> Enum.filter(&do_match(&1, base_list))
    |> Enum.map(&Enum.join/1)
  end

  defp do_match(l1, l2) do
    [ch | tail] = normalize(l1)
    normalized_l2 = normalize(l2)
    do_match(tail, normalized_l2 -- [ch], ch in normalized_l2)
  end

  defp do_match([ch | tail], l2, match?) when match? do
    do_match(tail, l2 -- [ch], ch in l2)
  end

  defp do_match([], [], match?) when match?, do: true
  defp do_match(_l1, _l2, _match), do: false

  defp normalize(list) do
    list
    |> Enum.map(&String.downcase/1)
  end

  defp case_match?(l1, l2) do
    Enum.join(l1) |> String.downcase() == Enum.join(l2) |> String.downcase()
  end
end

Anagram.match("nose", ~w(Eons ONES))
|> IO.inspect()
