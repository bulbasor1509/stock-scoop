import Wrapper from "@/components/Wrapper";
import Link from "next/link";


const Navbar = () => {
    return (
        <div className="top-0 sticky inset-x-0 w-full border-b border-gray-200 bg-white/75 backdrop-blur-lg transition-all duration-200">
            <Wrapper>
                <div className="text-lg flex justify-between items-center">
                    <div className="uppercase h-16 flex items-center">
                        stock scoop
                    </div>
                    <div className="uppercase h-16 flex items-center gap-4 text-sm">
                        <Link href="/bulk">bulk deals</Link>
                    </div>
                </div>
            </Wrapper>
        </div>
    )
}

export default Navbar