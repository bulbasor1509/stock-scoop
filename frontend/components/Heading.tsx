import {cn} from "@/lib/utils";

const Heading = ({heading, className}: {heading: string, className?: string}) => {
    return (
        <div className={cn("text-lg uppercase tracking-wide font-medium", className)}>
            {heading}
        </div>
    )
}

export default Heading