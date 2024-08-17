import { API_BASE_URL } from '#/common/constants';

export default function Home() {
    const items = fetch(`${API_BASE_URL}/items`)
        .then((response) => response.json())
        .then((data) => console.log(data));

    return (
        <>
            <h1>Menu</h1>
        </>
    );
}
