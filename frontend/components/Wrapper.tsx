import {cn} from "@/lib/utils";

const Wrapper = ({
    className,
    children,
}: {
    className?: string
    children: React.ReactNode
})  => {
    return (
        <div className={cn("mx-auto w-full px-6 md:px-12", className)}>
            {children}
        </div>
    )
}

export default Wrapper
