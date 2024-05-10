
import { SignedIn, SignedOut, RedirectToSignIn, UserButton } from "@clerk/nextjs";
import Link from 'next/link'
const Rankings = () => {

	return (
		<>
			<SignedIn>
				<div className="w-full h-[10%] bg-blue-600 p-3 flex justify-center">
					<div className="w-3/4 flex justify-between">
						<h1 className='text-white text-xl'>Promotion Rankings | Rankings</h1>
						<div className='text-white flex gap-4'>
							<Link href='/admin'>Home </Link>
							<Link href='/admin/athletes'>Athletes </Link>
							<Link href='/admin/rankings'>Rankings </Link>
						</div>
						<UserButton />
					</div>
				</div>
				<section className="w-full p-3 flex justify-center">
					<div className="w-3/4 flex justify-between">
						<h1>Rankings</h1></div>
				</section>
			</SignedIn>
			<SignedOut>
				<RedirectToSignIn />
			</SignedOut>
		</>
	);
}

export default Rankings
