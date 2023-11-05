defmodule FloydWarshall do
  def floyd_warshall(graph) when is_map(graph) do
    verticles = Map.keys(graph)
    n = length(verticles)
    n_range = 0..(n - 1)
    dist = initialize_distance_matrix(graph, verticles)

    n_range
    |> Enum.reduce(dist, fn v, acc -> handle_verticles(acc, v, n_range) end)
  end

  defp handle_verticles(dist, v, n_range) do
    for i <- n_range do
      for j <- n_range do
        update_distance(dist, v, i, j)
      end
    end
  end

  defp initialize_distance_matrix(graph, verticles) do
    n = length(verticles)

    for i <- 0..(n - 1) do
      for j <- 0..(n - 1) do
        if i == j, do: 0, else: Map.get(graph[Enum.at(verticles, i)], Enum.at(verticles, j), 999)
      end
    end
  end

  defp update_distance(dist, v, i, j) do
    if at(dist, i, j) > at(dist, i, v) + at(dist, v, j) do
      at(dist, i, v) + at(dist, v, j)
    else
      at(dist, i, j)
    end
  end

  defp at(dist, i, j) do
    get_in(dist, [Access.at(i), Access.at(j)])
  end
end

# Example usage with the graph from the previous example
graph = %{
  "A" => %{"A" => 0, "B" => 5, "C" => 999, "D" => 10},
  "B" => %{"A" => 999, "B" => 0, "C" => 3, "D" => 999},
  "C" => %{"A" => 999, "B" => 999, "C" => 0, "D" => 1},
  "D" => %{"A" => 999, "B" => 999, "C" => 999, "D" => 0}
}

result = FloydWarshall.floyd_warshall(graph)
IO.inspect(result)
