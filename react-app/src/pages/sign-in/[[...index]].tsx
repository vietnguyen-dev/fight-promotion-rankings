import { SignIn } from "@clerk/nextjs";

export default function Page() {
	return (
		<section className='h-screen flex justify-center pt-6 bg-blue-600'>
			<SignIn path="/sign-in" forceRedirectUrl="/admin" />
		</section>
	);
}
