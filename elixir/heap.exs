defmodule Heap.List do
  defstruct size: 0, data: nil

  def build_heap(arr) do
    map =
      arr
      |> Enum.with_index()
      |> Enum.reduce(%{}, fn {v, k}, acc -> Map.put(acc, k, v) end)

    heapify(%Heap.List{size: length(arr), data: map})
  end

  defp heapify(heap) do
    Map.values(heap.data)
    |> Enum.with_index()
    |> Enum.reduce(%{}, fn {v, k}, acc -> Map.put(acc, k, v) end)
    |> Map.keys()
    |> Enum.reduce(heap, &heapify/2)
  end

  defp heapify(i, %Heap.List{size: n, data: map} = acc) do
    root_min = i
    left = 2 * i + 1
    right = 2 * i + 2

    root_min = if root_min < n and map[root_min] > map[left], do: left, else: root_min
    root_min = if root_min < n and map[root_min] > map[right], do: right, else: root_min

    if root_min != i do
      tmp = map[i]

      map =
        Map.put(map, i, map[root_min])
        |> Map.put(root_min, tmp)

      heapify(root_min, %{acc | data: map})
    else
      acc
    end
  end

  def to_list(%Heap.List{size: 0}, acc), do: Enum.reverse(acc)

  def to_list(heap, acc) do
    {top, rest} = Map.pop(heap.data, Map.keys(heap.data) |> Enum.min())

    to_list(
      heapify(%{heap | data: rest, size: heap.size - 1}),
      [top | acc]
    )
  end
end

defmodule Heap do
  defstruct size: 0, data: nil

  def build_heap(arr) do
    build_heap(arr, %Heap{size: 0})
  end

  defp build_heap(arr, heap) do
    arr
    |> Enum.reduce(heap, &push/2)
  end

  defp push(e, %Heap{size: n, data: data}) do
    %Heap{size: n + 1, data: meld(data, {e, []})}
  end

  defp meld(nil, heap), do: heap
  defp meld(heap, nil), do: heap

  defp meld({k0, l0}, {k1, _} = r) when k0 < k1, do: {k0, [r | l0]}
  defp meld({_, _} = l, {k1, r0}), do: {k1, [l | r0]}

  def pop(%Heap{data: {_, heap}, size: n} = _heap),
    do: %Heap{data: heapify(heap), size: n - 1}

  def root(%Heap{data: {v, _}} = _heap), do: v
  def root(%Heap{data: nil, size: 0} = _heap), do: nil

  defp heapify([]), do: nil
  defp heapify([q]), do: q

  defp heapify([left, right | q]) do
    min = meld(left, right)
    meld(min, heapify(q))
  end

  def to_list(heap, acc) do
    case Heap.root(heap) do
      nil -> Enum.reverse(acc)
      root -> to_list(Heap.pop(heap), [root | acc])
    end
  end
end

heap_with_list =
  [1, 5, 6, 8, 9, 7, 3]
  |> Heap.List.build_heap()
  |> Heap.List.to_list([])
  |> IO.inspect()

heap =
  [1, 5, 6, 8, 9, 7, 3]
  |> Heap.build_heap()
  |> Heap.to_list([])
  |> IO.inspect()


(heap == heap_with_list) |> IO.inspect()
