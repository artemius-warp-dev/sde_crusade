defmodule UnionFind do
  def find_connected_components(n, edges) do
    parents = for i <- 0..(n - 1), into: %{}, do: {i, i}
    ranks = for i <- 0..(n - 1), into: %{}, do: {i, 1}
    state = {n, parents, ranks}

    Enum.reduce(edges, state, fn [n1, n2], acc ->
      {nodes, parents, ranks} = acc
      {res, updated_parents, upddated_ranks} = union(parents, ranks, n1, n2)
      IO.inspect(res, label: "Union result")
      IO.inspect(updated_parents)
      {nodes - res, updated_parents, upddated_ranks}
    end)
  end

  defp find(node, parents) do
    # IO.inspect({node, parents})
    # |> IO.inspect()
    potential_representative = parents[node]

    if potential_representative == node do
      node
    else
      find(potential_representative, parents)
    end
  end

  # defp find(node, paretns, parent_itself?), do: node
  # defp find(

  # and finally puts either one of the trees (representing the set) under the root node of the other tree. 
  defp union(parents, ranks, node_a, node_b) do
    IO.inspect({node_a, node_b})
    a_parent = find(node_a, parents) |> IO.inspect(label: "A")
    b_parent = find(node_b, parents) |> IO.inspect(label: "B")
    res = calculate(a_parent, b_parent)

    if ranks[a_parent] < ranks[b_parent] do
      {res, Map.put(parents, a_parent, b_parent),
       Map.update!(ranks, b_parent, &(&1 + ranks[a_parent]))}
    else
      {res, Map.put(parents, b_parent, a_parent),
       Map.update!(ranks, a_parent, &(&1 + ranks[b_parent]))}
    end
  end

  defp calculate(node, node), do: 0
  defp calculate(_, _), do: 1
end

n = 6
edges = [[0, 3], [1, 2]]

UnionFind.find_connected_components(n, edges)
|> IO.inspect()
