export default function SimplePage({ title }: { title: string }) {
  return (
    <div className="container" style={{ paddingTop: 80, paddingBottom: 80 }}>
      <h1 className="h2">{title}</h1>
    </div>
  );
}
