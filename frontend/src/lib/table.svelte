<script lang="ts">
  import {
    Button,
    Heading,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
  } from 'flowbite-svelte';
  import { type TableData } from "./table-data";

  let { name, data }: { name :string, data: TableData[] } = $props()
</script>
<div>
  <Heading tag="h3" class="mb-10">{name}</Heading>
  <Table>
    <TableHead>
      <TableHeadCell>Name</TableHeadCell>
      <TableHeadCell>Amount</TableHeadCell>
      <TableHeadCell>Date</TableHeadCell>
      <TableHeadCell>Received</TableHeadCell>
    </TableHead>
    <TableBody>
      {#each data as entry (entry.id)}
        <TableBodyRow>
          <TableBodyCell>{entry.name}</TableBodyCell>
          <TableBodyCell>{entry.amount}</TableBodyCell>
          <TableBodyCell>{entry.date.toLocaleDateString("en-AU")}</TableBodyCell>
          <TableBodyCell>{entry.received}</TableBodyCell>
        </TableBodyRow>
      {/each}
    </TableBody>
  </Table>
</div>

