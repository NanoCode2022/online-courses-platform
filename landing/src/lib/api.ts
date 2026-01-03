const API_URL = import.meta.env.PUBLIC_API_URL;

export async function getCourses() {
  const res = await fetch(`${API_URL}/courses`);
  return res.json();
}

export async function getCourse(id: string) {
  const res = await fetch(`${API_URL}/courses/${id}`);
  return res.json();
}
