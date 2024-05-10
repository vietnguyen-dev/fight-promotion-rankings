import { useState, useEffect } from 'react'
import { SignedIn, SignedOut, RedirectToSignIn, UserButton } from "@clerk/nextjs";
import Link from 'next/link'
import { getAuth, buildClerkProps, clerkClient } from "@clerk/nextjs/server";
import { GetServerSideProps } from "next";

import { iPromotions } from '../../interfaces/promotion'

interface iAdminProps {
	user: any
}

const Admin = (props: any) => {
	const [promotion, setPromotion] = useState<iPromotions | null>(null);

	useEffect(() => {
		const getPromotion = async (id: number) => {
			try {
				const res = await fetch(`http://localhost:8080/promotions/${id}`)
				const data = await res.json()
				console.log(data)
			}
			catch (err) {
				console.error(err)
			}

		}
		getPromotion(props.__clerk_ssr_state.user.publicMetadata.id);
	}, [])

	return (
		<>
			<SignedIn>
				<div className="w-full h-[10%] bg-blue-600 p-3 flex justify-center">
					<div className="w-3/4 flex justify-between">
						<h1 className='text-white text-xl'>Promotion Rankings</h1>
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


export const getServerSideProps: GetServerSideProps = async (ctx) => {
	const { userId } = getAuth(ctx.req);

	const user = userId ? await clerkClient.users.getUser(userId) : undefined;
	console.log(user, 'from server');

	let returnObject = { props: { ...buildClerkProps(ctx.req, { user }) } };
	return { props: { ...buildClerkProps(ctx.req, { user }) } };
};
