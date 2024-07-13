import { useState } from 'react';
import './TugasCRUDHooks.css';

const TugasCrudHooks = () => {
    const [daftarBuah, setDaftarBuah] = useState([
        { nama: "Nanas", hargaTotal: 100000, beratTotal: 4000 },
        { nama: "Manggis", hargaTotal: 350000, beratTotal: 10000 },
        { nama: "Nangka", hargaTotal: 90000, beratTotal: 2000 },
        { nama: "Durian", hargaTotal: 400000, beratTotal: 5000 },
        { nama: "Strawberry", hargaTotal: 120000, beratTotal: 6000 }
    ]);

    const [form, setForm] = useState({
        nama: "",
        hargaTotal: "",
        beratTotal: ""
    });

    const [editIndex, setEditIndex] = useState(null);

    const handleChange = (event) => {
        setForm({
            ...form,
            [event.target.name]: event.target.value
        });
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        
        if (form.nama && !isNaN(form.hargaTotal) && !isNaN(form.beratTotal) && form.beratTotal >= 2000) {
            if (editIndex === null) {
                setDaftarBuah([...daftarBuah, form]);
            } else {
                const newDaftarBuah = [...daftarBuah];
                newDaftarBuah[editIndex] = form;
                setDaftarBuah(newDaftarBuah);
                setEditIndex(null);
            }
            setForm({
                nama: "",
                hargaTotal: "",
                beratTotal: ""
            });
        } else {
            alert("Semua inputan wajib diisi, harga total dan berat total harus berupa angka, dan berat total minimal 2000 gram.");
        }
    };

    const handleEdit = (index) => {
        const buah = daftarBuah[index];
        setForm({
            nama: buah.nama,
            hargaTotal: buah.hargaTotal,
            beratTotal: buah.beratTotal
        });
        setEditIndex(index);
    };

    const handleDelete = (index) => {
        const newDaftarBuah = [...daftarBuah];
        newDaftarBuah.splice(index, 1);
        setDaftarBuah(newDaftarBuah);
    };

    const formatNumber = (num) => {
        return Number(num).toString();
    };

    const calculateHargaPerKg = (hargaTotal, beratTotal) => {
        return formatNumber((hargaTotal / (beratTotal / 1000)).toFixed(2));
    };

    return (
        <div className="tugas-crud-hooks">
            <h1>Tugas CRUD Hooks</h1>
            <form className="form" onSubmit={handleSubmit}>
                <label>
                    Nama:
                    <input type="text" name="nama" value={form.nama} onChange={handleChange} required />
                </label>
                <br />
                <label>
                    Harga Total:
                    <input type="number" name="hargaTotal" value={form.hargaTotal} onChange={handleChange} required />
                </label>
                <br />
                <label>
                    Berat Total (gram):
                    <input type="number" name="beratTotal" value={form.beratTotal} onChange={handleChange} required min="2000" />
                </label>
                <br />
                <button type="submit">Submit</button>
            </form>
            <table className="table">
                <thead>
                    <tr>
                        <th>No</th>
                        <th>Nama</th>
                        <th>Harga Total</th>
                        <th>Berat Total (kg)</th>
                        <th>Harga per Kg</th>
                        <th>Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    {daftarBuah.map((buah, index) => (
                        <tr key={index}>
                            <td>{index + 1}</td>
                            <td>{buah.nama}</td>
                            <td>{buah.hargaTotal}</td>
                            <td>{formatNumber((buah.beratTotal / 1000).toFixed(2))}</td>
                            <td>{calculateHargaPerKg(buah.hargaTotal, buah.beratTotal)}</td>
                            <td>
                                <button className="btn-edit" onClick={() => handleEdit(index)}>Edit</button>
                                <button className="btn-delete" onClick={() => handleDelete(index)}>Delete</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default TugasCrudHooks;
