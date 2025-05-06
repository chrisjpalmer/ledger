<script lang="ts">
    import Table from "$lib/table.svelte";
    import { type TableData } from "$lib/table-data";
    import IncomeTable from "$lib/incomeTable.svelte";
    import { Heading, Button } from "flowbite-svelte";
    import { LedgerApi, Configuration } from "$lib/api";

    let expenses:TableData[] = $state([
      {id: "1", name: "Fuel #1", amount:20.3, date: new Date("2024-12-01"), received: true},
      {id: "2", name: "Fuel #2", amount:20.3, date: new Date("2024-12-01"), received: true},
      {id: "3", name: "Fuel #3", amount:20.3, date: new Date("2024-12-01"), received: true},
    ])

    const cfg = new Configuration({
      basePath: "http://localhost:8080"
    });
    const api = new LedgerApi(cfg);

    
</script>
<style>
  .parent {
    height: 100%;
    padding: 10px;
    display: flex;
    /* grid-template-columns: 100%;
    #grid-template-rows: 50px 25% 75%; */
    row-gap: 20px;
    flex-direction: column;
  }
</style>
<div class="parent">
  <Heading tag="h1">Ledger</Heading>
  <IncomeTable api={api} month={0}></IncomeTable>
  <Table name="Expenses" data={expenses}></Table>
</div>
