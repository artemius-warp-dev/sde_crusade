defmodule Pov do
  @typedoc """
  A tree, which is made of a node with several branches
  """
  @type tree :: {any, [tree]}

  @doc """
  Reparent a tree on a selected node.
  """
  @spec from_pov(tree :: tree, node :: any) :: {:ok, tree} | {:error, atom}
  def from_pov(tree, node) do
    case dfs(tree, node) do
      nil -> {:error, :nonexistent_target}
      {pov_tree, visited} -> {:ok, change_parent(visited, pov_tree)}
    end
  end

  defp change_parent([], tree), do: tree

  defp change_parent([{parent_node, parent_subtree} | tl], {node, tree}) do
    {parent_node, List.keydelete(parent_subtree, node, 0)}
    |> then(&{node, [change_parent(tl, &1) | tree]})
  end

  defp dfs(tree, node, visited \\ [])
  defp dfs({node, _} = tree, node, visited), do: {tree, visited}

  defp dfs({node, tree}, target, visited) do
    tree
    |> Enum.reduce_while(nil, fn subtree, acc ->
      case dfs(subtree, target, [{node, tree} | visited]) do
        nil -> {:cont, acc}
        found -> {:halt, found}
      end
    end)
  end

  @doc """
  Finds a path between two nodes
  """
  @spec path_between(tree :: tree, from :: any, to :: any) :: {:ok, [any]} | {:error, atom}
  def path_between(tree, from, to) do
    with {:ok, pov_tree} <- from_pov(tree, from) |> IO.inspect(),
         {{parent, _}, visited} <- dfs(pov_tree, to) |> IO.inspect() do
      create_path(visited, [parent])
    else
      {:error, :nonexistent_target} -> {:error, :nonexistent_source}
      nil -> {:error, :nonexistent_destination}
    end
  end

  defp create_path([], acc), do: {:ok, acc}
  defp create_path([{node, _} | tail], acc), do: create_path(tail, [node | acc])
end

leaf = fn a -> {a, []} end

t =
  {:grandparent,
   [
     {:parent, [{:x, [leaf.(:kid0), leaf.(:kid1)]}, leaf.(:sibling0), leaf.(:sibling1)]},
     {:uncle, [leaf.(:cousin0), leaf.(:cousin1)]}
   ]}

# Pov.from_pov(t, :x) |> IO.inspect()

t = {:parent, [leaf.(:x), leaf.(:sibling)]}
Pov.path_between(t, :x, :parent) |> IO.inspect()
