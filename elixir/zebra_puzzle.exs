defmodule ZebraPuzzle do
  @moduledoc """
  Solve the zebra puzzle.
  """

  defmodule Solution do
    @nations [:englishman, :ukrainian, :spaniard, :norwegian, :japanese]
    @colors [:red, :green, :ivory, :yellow, :blue]
    @animals [:dog, :snails, :fox, :horse, :zebra]
    @drinks [:coffee, :tea, :milk, :juice, :water]
    @cigarets [:old_gold, :kools, :chesterfields, :lucky_strike, :parliaments]

    def solution() do
      for nations <- permutations(@nations),
          # 10
          index?(nations, :norwegian, 0) do
        for colors <- permutations(@colors),
            # 2
            matchof?(colors, :red, nations, :englishman),
            # 6
            leftof?(colors, :ivory, colors, :green),
            # 15
            nextof?(colors, :blue, nations, :norwegian) do
          for animals <- permutations(@animals),
              # 3
              matchof?(animals, :dog, nations, :spaniard) do
            for drinks <- permutations(@drinks),
                # 9
                index?(drinks, :milk, 2),
                # 4
                matchof?(drinks, :coffee, colors, :green),
                # 5
                matchof?(drinks, :tea, nations, :ukrainian) do
              for cigarets <- permutations(@cigarets),
                  # 7
                  matchof?(cigarets, :old_gold, animals, :snails),
                  # 8
                  matchof?(cigarets, :kools, colors, :yellow),
                  # 14
                  matchof?(cigarets, :parliaments, nations, :japanese),
                  # 13
                  matchof?(cigarets, :lucky_strike, drinks, :juice),
                  # 11
                  nextof?(cigarets, :chesterfields, animals, :fox),
                  # 12
                  nextof?(cigarets, :kools, animals, :horse) do
                List.zip([nations, colors, animals, drinks, cigarets])
              end
            end
          end
        end
      end
      |> List.flatten()
      |> Enum.reject(&(&1 == []))
    end

    defp index?(list, term, target) do
      Enum.find_index(list, &(&1 == term)) == target
    end

    defp matchof?(from, t1, target, t2) do
      i = Enum.find_index(from, &(&1 == t1))
      Enum.at(target, i) == t2
    end

    defp leftof?(from, t1, target, t2) do
      i = Enum.find_index(from, &(&1 == t1))
      Enum.at(target, i + 1) == t2
    end

    defp nextof?(from, t1, target, t2) do
      i = Enum.find_index(from, &(&1 == t1))
      Enum.at(target, i + 1) == t2 || (i > 0 && Enum.at(target, i - 1) == t2)
    end

    defp permutations([]), do: [[]]

    defp permutations(list),
      do: for(elem <- list, rest <- permutations(list -- [elem]), do: [elem | rest])
  end

  @solution Solution.solution()

  @doc """
  Determine who drinks the water
  """

  @spec drinks_water() :: atom
  def drinks_water() do
    @solution
    |> Enum.find_value(fn
      {nation, _, _, :water, _} -> nation
      _ -> false
    end)
  end

  @doc """
  Determine who owns the zebra
  """

  @spec owns_zebra() :: atom
  def owns_zebra() do
    @solution
    |> Enum.find_value(fn
      {nation, _, :zebra, _, _} -> nation
      _ -> false
    end)
  end
end

ZebraPuzzle.drinks_water()
