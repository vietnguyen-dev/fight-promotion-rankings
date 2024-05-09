import { SignedIn, SignedOut, RedirectToSignIn, UserButton } from "@clerk/nextjs";
const Admin = () => {

	return (
		<>
			<SignedIn>
				<div className="w-full h-[10%] bg-blue-600 p-3 flex justify-center">
					<div className="w-10/12 flex justify-between">
						<h1 className='text-white'>Promotion Rankings</h1>
						<UserButton />
					</div>
				</div>
				<section className="w-full p-3 flex justify-center">
					<div className="w-10/12 flex justify-between">
						<p>You are signed in</p>
					</div>
				</section>
			</SignedIn>
			<SignedOut>
				<RedirectToSignIn />
			</SignedOut>
		</>
	);
}

export default Admin
