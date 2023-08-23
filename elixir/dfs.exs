defmodule DFS do
  def dfs(graph, start_node) do
    result = []
    visited = []
    _dfs(graph, start_node, visited, result, visited?(start_node, visited))
    |> then(fn {visited, result} -> {Enum.reverse(visited), Enum.reverse(result)} end)
  end

  defp _dfs(graph, node, visited, result, visited?) when visited? do
    {visited, result}
  end

  defp _dfs(graph, node, visited, result, _) do
    new_visited = [node | visited]
    new_result = [node | result]
    neighbors = Map.get(graph, node, [])

    Enum.reduce(neighbors, {new_visited, new_result}, fn neighbor, {acc_visited, acc_result} ->
      {next_visited, next_result} =
        _dfs(graph, neighbor, acc_visited, acc_result, visited?(neighbor, acc_visited))

      {next_visited, next_result}
    end)
  end

  defp visited?(node, visited), do: node in visited
end

# Example graph represented as a adjency map

graph = %{
  "A" => ["B", "C"],
  "B" => ["D", "E"],
  "C" => ["F"],
  "D" => [],
  "E" => ["F"],
  "F" => []
}

# [A, B, D, E, F, C]

start_node = "A"
{visited_nodes, traversal_orders} = DFS.dfs(graph, start_node)

IO.puts("Visited nodes: #{inspect(visited_nodes)}")
IO.puts("Traversal order: #{inspect(traversal_orders)}")
