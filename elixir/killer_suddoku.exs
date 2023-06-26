defmodule KillerSudokuHelper do
  @doc """
  Return the possible combinations of `size` distinct numbers from 1-9 excluding `exclude` that sum up to `sum`.
  """
  @spec combinations(cage :: %{exclude: [integer], size: integer, sum: integer}) :: [[integer]]
  def combinations(cage) do
    for x <- 1..9,
        combination = combination(x, cage.sum, cage.size, []),
        Enum.all?(combination, fn x -> x > 0 end),
        not Enum.any?(combination, fn x -> x in cage.exclude end),
        length(Enum.uniq(combination)) == length(combination),
        Enum.sum(combination) == cage.sum do
      Enum.sort(combination)
    end
    |> IO.inspect()
    |> Enum.uniq()
  end

  defp combination(_, sum, 1, acc), do: [sum | acc]

  defp combination(x, sum, size, acc) do
    combination(x + 1, sum - x, size - 1, [x | acc])
  end
end

# @spec combinations(cage :: %{exclude: [integer], size: integer, sum: integer}) :: [[integer]]

# def combinations(cage) do
#   %{sum: sum, size: size, exclude: exclude} = cage

#   sorted_numbers = Enum.to_list(1..9) -- exclude

#   all_combinations = do_all_combinations(sorted_numbers, size)

#   Enum.filter(all_combinations, &(Enum.sum(&1) == sum))
# end

# defp do_all_combinations(sorted_numbers, 1), do: Enum.map(sorted_numbers, &[&1])

# defp do_all_combinations(sorted_numbers, size) do
#   Enum.flat_map(sorted_numbers, fn n ->
#     sorted_numbers
#     |> Enum.filter(&(&1 > n))
#     |> do_all_combinations(size - 1)
#     |> Enum.map(&[n | &1])
#   end)
# end

# cage = %{exclude: [], size: 2, sum: 10}
# cage = %{exclude: [], size: 9, sum: 45}
# cage = %{exclude: [], size: 3, sum: 7}
cage = %{exclude: [1, 4], size: 2, sum: 10}
# == [[1, 9], [2, 8], [3, 7], [4, 6]]
KillerSudokuHelper.combinations(cage) |> IO.inspect()
