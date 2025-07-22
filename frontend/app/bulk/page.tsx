"use client"
import DataTable from "@/components/Table";
import {ColumnDef} from "@tanstack/table-core";
import {Bulk, BulkDataResponse} from "@/types/bulk";
import Heading from "@/components/Heading";
import Wrapper from "@/components/Wrapper";
import {useEffect, useState} from "react";

const bulk_table: ColumnDef<Bulk>[] = [
    {
        accessorKey: "BD_DT_DATE",
        header: "Date",
        cell: ({getValue}) => {
            const value = getValue<string>()
            return value.split("T")[0]
        }
    },
    {
        accessorKey: "BD_SYMBOL",
        header: "Symbol",
    },
    {
        accessorKey: "BD_SCRIP_NAME",
        header: "Security",
    },
    {
        accessorKey: "BD_CLIENT_NAME",
        header: "Client",
    },
    {
        accessorKey: "BD_BUY_SELL",
        header: "Type",
    },
    {
        accessorKey: "BD_QTY_TRD",
        header: "Quantity",
    },
    {
        accessorKey: "BD_TP_WATP",
        header: "Price",
        // cell: ({getValue}) => {
        //     const value = getValue<string>()
        //     return parseFloat(value)
        // }
    }
]

const Home =  () => {
    const [data, setData] = useState<Bulk[]>([])
    const [isLoading, setLoading] = useState(true)

    useEffect(() => {
        fetch("http://localhost:8080/bd_trade")
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