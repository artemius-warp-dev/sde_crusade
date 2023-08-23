defmodule BFS do
  def bfs(graph, start_node) do
    result = []
    visited = []

    _bfs(graph, [start_node], visited, result, visited?([start_node], visited))
    |> then(fn {visited, result} -> {Enum.reverse(visited), result} end)
  end

  defp _bfs(graph, [node | rest], visited, result, visited?) when visited? do
    {visited, result}
  end

  defp _bfs(graph, [node | rest], visited, result, _) do
    new_visited = [node | visited]
    new_result = result ++ [node]

    neighbors = Map.get(graph, node, [])
    new_nodes = rest ++ neighbors

    _bfs(
      graph,
      new_nodes,
      new_visited,
      new_result,
      visited?(new_nodes, new_visited)
    )
  end

  defp visited?([node | _], visited), do: node in visited
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

# [A, B, C, D, E, F]

start_node = "A"
{visited_nodes, traversal_orders} = BFS.bfs(graph, start_node)

IO.puts("Visited nodes: #{inspect(visited_nodes)}")
IO.puts("Traversal order: #{inspect(traversal_orders)}")
