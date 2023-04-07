defmodule BinarySearch do
  @doc """
    Searches for a key in the tuple using the binary search algorithm.
    It returns :not_found if the key is not in the tuple.
    Otherwise returns {:ok, index}.

    ## Examples

      iex> BinarySearch.search({}, 2)
      :not_found

      iex> BinarySearch.search({1, 3, 5}, 2)
      :not_found

      iex> BinarySearch.search({1, 3, 5}, 5)
      {:ok, 2}

  """

  @spec search(tuple, integer) :: {:ok, integer} | :not_found
  def search(numbers, key) do
    search(numbers, key, 0, tuple_size(numbers) - 1)
  end

  defp search({}, _, _, _), do: :not_found
  defp search(_, _, left, right) when right < left, do: :not_found

  defp search(numbers, key, left, right) do
    mid = div(left + right, 2)
    mid_elem = elem(numbers, mid)

    cond do
      mid_elem == key -> {:ok, mid}
      mid_elem < key -> search(numbers, key, mid + 1, right)
      mid_elem > key -> search(numbers, key, left, mid - 1)
    end
  end
end

BinarySearch.search({1, 3, 4, 5, 8, 9, 11}, 1) |> IO.inspect()
