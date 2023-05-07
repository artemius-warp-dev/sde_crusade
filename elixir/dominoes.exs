defmodule Dominoes do
  @type domino :: {1..6, 1..6}

  @doc """
  chain?/1 takes a list of domino stones and returns boolean indicating if it's
  possible to make a full chain
  """
  @spec chain?(dominoes :: [domino]) :: boolean
  def chain?([]), do: true

  def chain?(dominoes) do
    dominoes
    |> normalize()
    |> construct(length(dominoes))
    # {f1, l1} = List.first(dominoes)
    # {f2, l2} = List.last(dominoes)
    # f1 == l2
  end

  # [{1, 2}, {3, 1}, {2, 3}]

  defp normalize(dominoes) do
    dominoes
    |> Enum.reduce([], {h,t}, acc -> [{h,t}] ++ [{t,h}] ++ acc  )
  end


  defp construct([], _, _, _), do: false
  defp consturct(_, 0, solution, acc), do: [solution | acc]
  defp construct([{h,t} | tail], length) do
    case Enum.find(tail, fn {x, _} x == t end) do
      nil -> construct(tail, length)
      domino -> 
    end
  end

 

  

  #   {1, 2}, {2, 3}, {3, 1}
  #    %{
  #      {1,2} => [{2,3}, {3,1}]
  #      {2, 1} => {1,3}, {1,2}
  #      {3, 1} => {1,2}
  #      {1, 3} => {3,2}
  #      {2, 3} => {3, 1}
  #      {3, 2} => {2, 1}
  # }
end

Dominoes.format({1, 2}, [{1, 2}, {3, 1}, {2, 3}]) |> IO.inspect()
