defmodule TopologicalSort do
  def sort({graph, n}) do
    state = %{graph: graph, visited: %{}, result: []}

    0..(n - 1)
    |> Enum.reverse()
    |> Enum.reduce(state, &handle_verticle/2)
    |> then(& &1.result)
  end

  defp handle_verticle(node, %{visited: visited} = state)
       when is_map_key(visited, node),
       do: state

  defp handle_verticle(node, state), do: dfs(node, state)

  defp dfs(node, state) do
    updated_visited = Map.put(state.visited, node, true)
    neighbors = Map.get(state.graph, node, [])
    # usual dfs order, you should reverse result upper
    #updated_state = %{state | visited: updated_visited, result: [node | state.result]}
    # topological order
    updated_state = %{state | visited: updated_visited}

    Enum.reduce(neighbors, updated_state, &handle_neighbor/2)
    |> then(&%{&1 | result: [node | &1.result]})
  end

  defp handle_neighbor(node, %{visited: visited} = state) when is_map_key(visited, node),
    do: state

  defp handle_neighbor(node, state), do: dfs(node, state)
end

graph =
  [{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}]
  |> Enum.reduce(%{}, fn {v1, v2}, acc ->
    Map.update(acc, v1, [v2], &[v2 | &1])
  end)

TopologicalSort.sort({graph, 6})
|> IO.inspect(label: "RESULT")
