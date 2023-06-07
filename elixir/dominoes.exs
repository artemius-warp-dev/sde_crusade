defmodule Dominoes do
  @type domino :: {1..6, 1..6}

  @doc """
  chain?/1 takes a list of domino stones and returns boolean indicating if it's
  possible to make a full chain
  """
  @spec chain?(dominoes :: [domino]) :: boolean
  def chain?([]), do: true
  def chain?([{f,f}]), do: true
  def chain?(dominoes) do
    dominoes
    |> Enum.any?(fn {f,s} = bone ->
      updated_rest = List.delete(dominoes, bone)
      do_chain({f,s}, {f,s}, updated_rest) or
      do_chain({s,f}, {s,f}, updated_rest)
    end)
  end

 
  defp do_chain({f, _s}, _, [{_lf, f}]) do
    true
  end

  defp do_chain(_, _, []), do: false

  defp do_chain(first, {f, s}, rest) do
    case get_next(rest, {f, s}) do
      nil ->
        false

      {nf, ns} = next ->
        do_chain(first, {nf, ns}, List.delete(rest, next)) or
          do_chain(first, {ns, nf}, List.delete(rest, next))
    end
  end

  defp get_next(dominoes, {first, second}) do
    dominoes
    |> Enum.find(fn {f, s} -> f == second or s == first end)
  end
end

# Dominoes.chain?([{2, 1}, {2, 3}, {3, 1}])
# Dominoes.chain?([{1, 2}, {2, 3}, {3, 1}, {4, 5}, {5, 6}, {6, 4}])
# Dominoes.chain?([{1, 2}, {2, 1}, {3, 4}, {4, 3}])
 #Dominoes.chain?([{1, 2}, {2, 3}, {3, 1}, {2, 4}, {2, 4}])
#Dominoes.chain?([{1, 2}, {2, 3}, {3, 4}])
# Dominoes.chain?([{1, 2}])
#Dominoes.chain?([{1, 1}])
Dominoes.chain?([{1, 2}, {1, 3}, {2, 3}])
|> IO.inspect()
