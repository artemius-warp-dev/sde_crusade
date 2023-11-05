defmodule Prim do
  def minimum_spanning_tree(graph) when is_map(graph) do
    vertex = Map.keys(graph) |> List.first()
    minimum_spanning_tree(graph, [vertex], [])
  end

  defp minimum_spanning_tree(graph, visited, edges) when length(visited) == map_size(graph),
    do: edges

  defp minimum_spanning_tree(graph, visited, edges) do
    new_edge = choose_next_edge(graph, visited)
    {_v1, v2, _} = new_edge
    minimum_spanning_tree(graph, [v2 | visited], [new_edge | edges])
  end

  defp choose_next_edge(graph, visited) do
    potential_edges =
      for v <- visited,
          u <- Map.keys(graph[v]),
          u not in visited,
          do: {v, u, graph[v][u]}

    Enum.min_by(potential_edges, fn {_, _, weight} -> weight end)
  end
end

# Example usage
graph = %{
  "A" => %{"B" => 4, "C" => 1},
  "B" => %{"A" => 4, "C" => 2, "D" => 3, "E" => 2},
  "C" => %{"A" => 1, "B" => 2, "D" => 4},
  "D" => %{"B" => 3, "C" => 4, "E" => 3},
  "E" => %{"B" => 2, "D" => 3}
}

graph
|> Prim.minimum_spanning_tree()
|> IO.inspect()
