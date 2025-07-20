"use client"
import DataTable from "@/components/Table";
import {ColumnDef} from "@tanstack/table-core";
import {Bulk, BulkDataResponse} from "@/types/bulk";
import Heading from "@/components/Heading";
import Wrapper from "@/components/Wrapper";
import {useEffect, useState} from "react";

const bulk_table: ColumnDef<Bulk>[] = [
    {
        accessorKey: "date",
        header: "Date",
        cell: ({getValue}) => {
            const value = getValue<string>()
            return value.split("T")[0]
        }
    },
    {
        accessorKey: "symbol",
        header: "Symbol",
    },
    {
        accessorKey: "security",
        header: "Security",
    },
    {
        accessorKey: "client",
        header: "Client",
    },
    {
        accessorKey: "type",
        header: "Type",
    },
    {
        accessorKey: "trade_qty",
        header: "Quantity",
    },
    {
        accessorKey: "avg_price",
        header: "Price",
        cell: ({getValue}) => {
            const value = getValue<string>()
            return parseFloat(value)
        }
    }
]

const Home =  () => {
    const [data, setData] = useState<Bulk[]>([])
    const [isLoading, setLoading] = useState(true)

    useEffect(() => {
        fetch("http://localhost:8080/bulk")
            .then((res) => res.json())
            .then((data: BulkDataResponse) => {
                setData(data.data)
                setLoading(false)
            })
    }, [])


    return (
        <>
            <Wrapper>
                <Heading
                    heading="bulk deals historical data"
                    className="mt-8"
                />
                <div className="py-4">
                    <DataTable columns={bulk_table} data={data}/>
                </div>
            </Wrapper>
        </>
    )
}

export default Home