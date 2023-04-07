defmodule Domain.MoneyTransferService do
  @spec transfer(user, user, amount) :: {:ok, transfer_id} || {:error, reason}
  def transfer(user1, user2, amount) do
    [user1, user2, amount]
    |> Transaction.start(Bank.get_money() / 2)
    |> Transaction.add(Bank.put_money() / 2, Bank.get_status() / 1)
    |> Transaction.resolve()
  end
end

defmodule Transaction do
  def on_success(_id, res), do: {:ok, res}
  def on_error(id, reason) do
    rollback(id)
    {:error, reason}
  end

  @spec start(keyword, function) :: {:ok, transaction_id}
  def start(args, callback) do
    {:ok, id} = callback.(args)
    {:ok, id}
  end

  def add({:ok, id} = res, callback, check_callback) do
    case poll(id, check_callback) do
      {:ok, args} ->
        callback.(args)

      {:error, reason} ->
        {:error, reason}
    end
  end

  def add({:error, reason}, _, _), do: {:error, reason}

  defp poll(id, callback) do
    case callback.(id) do
      {:ok, :pending} ->
        poll(id, callback)

      {:error, reason} ->
        on_error(id, reason)

      {:ok, result} ->
        on_success(id, result) 
    end
  end

  def rollback(id), do: :rollback_transaction

  def resolve({:ok, id}), do: {:ok, id}
  def resolve({:error, reason}), do: {:error, resolve}
end

defmodule Bank do
  def get_money(user, amount) do
    Process.sleep(10_000)
    {:ok, id}
  end

  def get_status(id), do: {:ok, :pending}

  def put_money(user, amount) do
    Process.sleep(20_000)
    {:ok, id}
  end
end

# defmodule Domain.MoneyTransferService do

#   def transfer(%User{} = user1, %User{} = user2, amount) do
#     user1
#     |> Transfer.create(amount)
#     |> Transfer.put_to(user2)
#   end
# end

# defmodule TransferPromise do

# 	@callback ....

#   # redis
# end

# defmodule Transfer do

#   defstruct [
#   	:from,
#     :to,
#     :amount,
#     :from_transaction,
#     :to_transaction,
#     :promise
#   ]

#   def create(user, amount) do

#     {:ok, transaction_id, transaction_promise} = Bank.get_money(user, amount,
#     	on_success: from_transaction_success/1,
#       on_fail: from_transaction_failed/1
#     )

#     %Transfer{
#     	from: user,
#       amount: amount,
#       from_transaction: transaction_id,
#       promise: transaction_promise
#     }
#   end

#   def from_transaction_success(transaction) do

#   	{:ok, transaction_id, transaction_promise} = Bank.put_money(user, amount,
#     	on_success: to_transaction_success/1,
#       on_fail: to_transaction_failed/1
#     )
#   end

#   def to_transaction_failed(transfer) do

#   	#{:ok, transaction_id, transaction_promise} = Bank.put_money(transfer.from, transfer.amount) 
#     Bank.cancel_transacation(transfer.from_transaction)

#     transfer
#   end
# end

# defmodule Bank do
# 	def get_money(user1, amount) do

# 	 	{:ok, id} | :error

#     # Process.sleep(10_000)
#   end

#   def put_money(user2, amount) do
#     Process.sleep(10_000)

#     :ok | :error
#   end
# end
