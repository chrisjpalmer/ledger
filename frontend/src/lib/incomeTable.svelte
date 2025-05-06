<script lang="ts">
  import {
    Button,
    Checkbox,
    Heading,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
    Input,
  } from 'flowbite-svelte';
  import { PlusOutline } from 'flowbite-svelte-icons';
  import type { Income, LedgerApi } from './api';

  let { api, month }: { api:LedgerApi, month: number } = $props()
  let rows:Income[] = $state([]);
  let newRows:Income[] = $state([]);
  let newRowCounter = 0;

  async function saveIncome(idx: number) {
    const incRow = newRows.splice(idx, 1)[0]
    await api.addIncome({income:incRow, month: month});
    await getIncome();
  }

  function addEmptyIncomeRow() {
    newRows.push({name:"hello", amount: 0, id: `${newRowCounter}`, date: new Date(), received: false})
    newRowCounter++;
  }

  async function getIncome() {
    const rs = await api.getIncome({month: month})
    rows = rs.income;
  }

  getIncome()
</script>
<div>
  <Heading tag="h3" class="mb-10">Income</Heading>
  <Table>
    <TableHead>
      <TableHeadCell>Name</TableHeadCell>
      <TableHeadCell>Amount</TableHeadCell>
      <TableHeadCell>Date</TableHeadCell>
      <TableHeadCell>Received</TableHeadCell>
      <TableHeadCell>-</TableHeadCell>
    </TableHead>
    <TableBody>
      {#each rows as entry, index (entry.id)} <!-- eventually add .id -->
        <TableBodyRow>
          <TableBodyCell>{entry.name}</TableBodyCell>
          <TableBodyCell>{entry.amount}</TableBodyCell>
          <TableBodyCell>{entry.date.toLocaleDateString("en-AU")}</TableBodyCell>
          <TableBodyCell>{entry.received}</TableBodyCell>
          <TableBodyCell></TableBodyCell>
        </TableBodyRow>
      {/each}
      {#each newRows as entry, index (entry.id)} <!-- eventually add .id -->
      <TableBodyRow>
        <TableBodyCell><Input bind:value={newRows[index].name} type="text" id="name" placeholder="Monthly Pay" required /></TableBodyCell>
        <TableBodyCell><Input bind:value={newRows[index].amount} type="number" id="amount" placeholder="22.00" required /></TableBodyCell>
          <TableBodyCell><Input bind:value={() => newRows[index].date.toISOString().split('T')[0], (v) => newRows[index].date = new Date(v)} type="date" id="date" placeholder="2025-05-04" required /></TableBodyCell>
        <TableBodyCell><Checkbox bind:checked={newRows[index].received} id="received" /></TableBodyCell>
          <TableBodyCell><Button onclick={() => saveIncome(index)}>Save</Button></TableBodyCell>
      </TableBodyRow>
      {/each}
      <TableBodyRow>
        <TableBodyCell></TableBodyCell>
        <TableBodyCell></TableBodyCell>
        <TableBodyCell></TableBodyCell>
        <TableBodyCell></TableBodyCell>
        <TableBodyCell>
          <Button onclick={addEmptyIncomeRow} pill={true} outline={true} class="p-2!" size="xl">
            <PlusOutline class="w-3 h-3 text-primary-700" />
          </Button>
        </TableBodyCell>
      </TableBodyRow>
    </TableBody>
  </Table>
</div>

