defmodule MinHeap do
  defstruct data: []

  @doc """
  Creates a new empty heap.
  """
  def new(data), do: %MinHeap{data: data}

  @doc """
  Inserts a value into the heap.
  """
  def insert(heap, value) do
    # IO.inspect({heap, value})
    heap = %{heap | data: heap.data ++ [value]}
    # |> IO.inspect(label: "UPHEAPIFY")
    up_heapify(heap, length(heap.data) - 1)
  end

  @doc """
  Pops the minimum value from the heap.
  """
  def pop(%MinHeap{data: []} = heap), do: {nil, heap}

  def pop(heap) do
    root_value = hd(heap.data)
    rest_data = tl(heap.data)
    last_index = length(rest_data) - 1
    swapped_heap = swap_elements(%{heap | data: rest_data}, 0, last_index)
    # IO.inspect(swapped_heap, label: "SWAPPED")
    {root_value, down_heapify(swapped_heap, 0)}
  end

  # Restores the heap property by moving an element up to its correct position.
  defp up_heapify(heap, 0), do: heap

  defp up_heapify(heap, index) when index > 0 do
    parent_index = div(index - 1, 2)

    if fetch_weight(index, heap.data) < fetch_weight(parent_index, heap.data) do
      swapped_heap = swap_elements(heap, index, parent_index)
      up_heapify(swapped_heap, parent_index)
    else
      heap
    end
  end

  # Restores the heap property by moving an element down to its correct position.
  defp down_heapify(heap, index) when index >= length(heap.data), do: heap

  defp down_heapify(heap, index) do
    left_child_index = index * 2 + 1
    right_child_index = index * 2 + 2

    min_child_index = get_min_index(heap, left_child_index, right_child_index)
    # IO.inspect(heap)

    if min_child_index < length(heap.data) and
         fetch_weight(min_child_index, heap.data) < fetch_weight(index, heap.data) do
      swapped_heap = swap_elements(heap, index, min_child_index)
      down_heapify(swapped_heap, min_child_index)
    else
      heap
    end
  end

  # Determines the index of the smallest child of a given node.
  defp get_min_index(heap, left_child_index, right_child_index) do
    if right_child_index < length(heap.data) and
         fetch_weight(right_child_index, heap.data) < fetch_weight(left_child_index, heap.data) do
      right_child_index
    else
      left_child_index
    end
  end

  # Swaps elements at two specified indices in the heap data list.
  defp swap_elements(heap, _, -1), do: heap

  defp swap_elements(heap, index1, index2) do
    data = heap.data
    temp = elem_at(index1, data)
    data = List.replace_at(data, index1, elem_at(index2, data))
    data = List.replace_at(data, index2, temp)
    %{heap | data: data}
  end

  # Retrieves an element from a specific index in the list.
  defp elem_at(index, list) when index >= 0 and index < length(list),
    do: Enum.at(list, index)

  defp fetch_weight(index, list), do: elem_at(index, list) |> elem(0)
end

defmodule DikjstraImp do
  # TODO
  # Ask question on elixirforum about:  
  # ** (ArgumentError) invalid right argument for operator "in", it expects a compile-time proper list or compile-time range on the right side when used in guard expressions, got: visited
  #     (elixir 1.14.2) lib/kernel.ex:4342: Kernel.raise_on_invalid_args_in_2/1
  #     (elixir 1.14.2) expanding macro: Kernel.in/2
  #     dikjstra_implementation.exs:134: DikjstraImp.bfs/2

  def init_graph(init_list, nodes) do
    init_graph = for verticle <- 1..nodes, into: %{}, do: {verticle, []}

    edges = for [u, v, w] <- init_list, do: {u, {v, w}}

    Enum.reduce(edges, init_graph, fn {u, {v, w}}, acc ->
      Map.update(acc, u, [{v, w}], fn edges -> [{v, w} | edges] end)
    end)
  end

  def find_shortest_path(graph, start_node) do
    visited = %{}
    minHeap = MinHeap.new([{0, start_node}])
    time = 0
    state = %{graph: graph, heap: minHeap, visited: visited, result: time}
    IO.inspect(graph, label: "GRAPH")

    find_shortest_path(state)
    |> then(& &1.result)
  end

  defp find_shortest_path(%{heap: %MinHeap{data: []}} = state), do: state

  defp find_shortest_path(state) do
    {{weight, node}, updated_heap} = MinHeap.pop(state.heap)

    updated_state = %{state | heap: updated_heap}

    bfs({weight, node}, updated_state)
    |> find_shortest_path()
  end

  defp bfs({_, node}, %{visited: visited} = state) when is_map_key(visited, node),
    do: find_shortest_path(state)

  defp bfs({weight, node} = parent, state) do
    neighbors = Map.get(state.graph, node, [])

    updated_state =
      %{
        state
        | visited: Map.put(state.visited, node, true),
          result: Enum.max([state.result, weight])
      }
      |> Map.put(:parent, %{node: node, weight: weight})

    neighbors
    |> Enum.reduce(updated_state, &handle_neighbor/2)
  end

  defp handle_neighbor({node, _weight}, %{visited: visited} = state)
       when is_map_key(visited, node),
       do: state

  defp handle_neighbor({node, weight}, state) do
    updated_heap = MinHeap.insert(state.heap, {weight + state.parent.weight, node})

    %{state | heap: updated_heap}
  end
end

# leetcode 743
# You are given a network of n nodes, labeled from 1 to n. You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.

# We will send a signal from a given node k. Return the minimum time it takes for all the n nodes to receive the signal. If it is impossible for all the n nodes to receive the signal, return -1.

times = [[2, 1, 1], [2, 3, 1], [3, 4, 1], [3, 5, 3], [4, 5, 3]]
nodes = 5
k = 2

DikjstraImp.init_graph(times, nodes)
|> DikjstraImp.find_shortest_path(k)
|> IO.inspect()
