import {ColumnDef} from "@tanstack/table-core";

export interface TableProps<TData, TValue> {
    columns: ColumnDef<TData, TValue>[]
    data: TData[]
}

