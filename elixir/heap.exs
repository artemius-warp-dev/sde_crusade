defmodule MinHeap do
  @moduledoc """
  The `MinHeap` module provides an implementation of a min-heap data structure.
  It allows you to create a heap, insert values, build a heap from a list, and pop the minimum value from the heap.
  """

  defstruct data: []

  @doc """
  Creates a new empty heap.
  """
  def new, do: %MinHeap{}
  def new(data), do: %MinHeap{data: data}

  @doc """
  Inserts a value into the heap.
  """
  def insert(heap, value) when is_number(value) do
    heap = %{heap | data: heap.data ++ [value]}
    up_heapify(heap, length(heap.data) - 1)
  end

  @doc """
  Builds a heap from a list of values.
  """
  def build_heap(list) do
    list
    |> Enum.reduce(%MinHeap{}, fn e, acc -> insert(acc, e) end)
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
    {root_value, down_heapify(swapped_heap, 0)}
  end

  @doc """
  Converts the heap to a list.
  """
  def to_list(heap, acc) do
    case pop(heap) do
      {nil, _} -> Enum.reverse(acc)
      {root, updated_heap} -> to_list(updated_heap, [root | acc])
    end
  end

  # Retrieves an element from a specific index in the list.
  defp elem_at(index, list) when index >= 0 and index < length(list), do: Enum.at(list, index)

  # Restores the heap property by moving an element up to its correct position.
  defp up_heapify(heap, 0), do: heap

  defp up_heapify(heap, index) when index > 0 do
    parent_index = div(index - 1, 2)

    if elem_at(index, heap.data) < elem_at(parent_index, heap.data) do
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

    if min_child_index < length(heap.data) and
         elem_at(min_child_index, heap.data) < elem_at(index, heap.data) do
      swapped_heap = swap_elements(heap, index, min_child_index)
      down_heapify(swapped_heap, min_child_index)
    else
      heap
    end
  end

  # Determines the index of the smallest child of a given node.
  defp get_min_index(heap, left_child_index, right_child_index) do
    if right_child_index < length(heap.data) and
         elem_at(right_child_index, heap.data) < elem_at(left_child_index, heap.data) do
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
end

# # Example usage of the MinHeap module...

# heap = MinHeap.new()
# heap = MinHeap.insert(heap, 5)
# heap = MinHeap.insert(heap, 300)
# heap = MinHeap.insert(heap, 3)
# heap = MinHeap.insert(heap, 100)
# heap = MinHeap.insert(heap, 8)

# {min_value, heap} = MinHeap.pop(heap)
# # Output: Minimum value: 3
# IO.puts("Minimum value: #{min_value}")

# heap =
#   [1, 5, 6, 8, 9, 7, 3]
#   |> MinHeap.build_heap()
#   |> MinHeap.to_list([])
#   |> IO.inspect()
