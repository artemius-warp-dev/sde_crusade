defmodule Alphametics do
  @type puzzle :: binary
  @type solution :: %{required(?A..?Z) => 0..9}

  @doc """
  Takes an alphametics puzzle and returns a solution where every letter
  replaced by its number will make a valid equation. Returns `nil` when
  there is no valid solution to the given puzzle.

  ## Examples

    iex> Alphametics.solve("I + BB == ILL")
    %{?I => 1, ?B => 9, ?L => 0}

    iex> Alphametics.solve("A == B")
    nil
  """
  @spec solve(puzzle) :: solution | nil
  def solve(puzzle) do
    # %{
    #   ?T => 0,
    #   ?O => 0,
    #   ?G => 0
    # }
    result = ["M", "O", "N", "E", "Y"]
    left = ["S", "E", "N", "D"]
    right = ["M", "O", "R", "E"]
    max_sum = get_max_sum(left, right)

    acc = %{"M" => 1}
    solve(left, right, result, max_sum, acc)
  end

  defp parse(), do: []

  defp solve(left, right, result, max_sum, acc) do
    pleft =
      1000..9999
      |> Enum.reduce([], fn x, acc ->
        cond do
          Integer.digits(x) |> Enum.uniq() |> length() == 4 ->
            [x | acc]

          true ->
            acc
        end
      end)

    pright =
      1000..9999
      |> Enum.reduce([], fn x, acc ->
        cond do
          Integer.digits(x) |> Enum.uniq() |> length() == 4 ->
            [x | acc]

          true ->
            acc
        end
      end)

    for x <- pleft,
        y <- pright,
        sum = x + y,
        sum >= 10000,
                    [s1, s2, s3, s4, s5] = sum_digits = Integer.digits(sum),
        Enum.uniq(sum_digits) |> length() == 5, 
        [l1, l2, l3, l4] = Integer.digits(x),
        [r1, r2, r3, r4] = Integer.digits(y),
        l2 == r4 and r4 == s4 and l3 == s3 and r2 == s2 and r1 == s1 and r1 == 1,
        l3 != r3 and l1 != r3 and l4 != r3 do
        
      [x, y]
    end
  end

  defp solve_reminder(acc, left, right, result) do
    # 100..max_sum #FIXME
    # |> Enum.reduce(0, )

    # [0..9]1 + [0..9]1 = 1[0..9][0..9]
    # x1 + y1 = 1zx
    for x <- 0..9, y <- 0..9 do
      # for y <- 0..9  do
      ["#{x}1", "#{y}1"] |> Enum.map(&String.to_integer/1)
      # end
    end
    # |> List.flatten()
    # |> Enum.chunk_every(2)
    |> Enum.filter(fn x -> Enum.sum(x) >= 100 end)
    |> Enum.find(fn [x, y] ->
      [l1, l2] = Integer.digits(x)
      [r1, r2] = Integer.digits(y)
      [d1, d2, d3] = (x + y) |> Integer.digits()
      l1 == d3 and d2 not in [l1, l2, r1, r2]
    end)
  end

  # defp brute_solve() do
  #   cond do
  #   end
  # end

  defp solve_most_left(left, right, result, acc) do
    # condition?
    Map.put(acc, result |> hd, 1)
  end

  # FIXME
  defp get_max_sum(left, right), do: 19998
end

Alphametics.solve("SEND + MORE == MONEY") |> IO.inspect()
# 21 + 81 == 102
# 21 + 91 == 112
